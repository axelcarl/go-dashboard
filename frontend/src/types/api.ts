import * as z from "zod";

export const Payment = z.object({
  id: z.number(),
  sender: z.string(),
  recipient: z.string(),
  amount: z.number(),
  createdAt: z.date(),
  updatedAt: z.date(),
});
