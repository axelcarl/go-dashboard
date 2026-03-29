import { api } from "@/lib/api-client";
import { type Payment } from "@/types/api";
import { useMutation, useQueryClient } from "@tanstack/react-query";

async function updatePayment(payment: Payment): Promise<void> {
  await api.put(`/payments/${payment.id}`, {
    amount: payment.amount,
    sender: payment.sender,
    recipient: payment.recipient,
  });
}

export function useUpdatePayment() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: updatePayment,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["payments"] });
    },
  });
}
